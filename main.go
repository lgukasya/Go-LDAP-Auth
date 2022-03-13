// Simple Go LDAP App

package main

import (
  "fmt"
  "log"
  
  "github.com/go-ldap/ldap/v3"
)

const (
    username = "cn=admin,dc=scytalelabs,dc=com"
    password = "adminPassword"
    Filter   = "(cn=*)"
    BaseDN   = "dc=scytalelabs,dc=com"
    Host     = "ldap://localhost"
)

// Connection ping
func Connect() (*ldap.Conn, error) {
    l, err := ldap.DialURL(fmt.Sprintf("%s:389", Host))
    return l, err
}

func Authentication(l *ldap.Conn) (error) {
    err := l.Bind(username, password);
    return err
}

// Query
func SearchRequest(l *ldap.Conn) (*ldap.SearchResult, error) {
    searchReq := ldap.NewSearchRequest(
        BaseDN,
        ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
        0,0, 
        false, Filter, 
        []string{},
        nil,
    )
    result, err := l.Search(searchReq)
    
    if err != nil {
        return nil, fmt.Errorf("Search Error: %s", err)
    }

    if len(result.Entries) > 0 {
        return result, nil
    } else {
        return nil, fmt.Errorf("Couldn't fetch search entries")
    }
}

// Main
func main() {
    // Connection
    l, err := Connect()
    
    if err != nil { 
        log.Fatal(err) 
    }
    
    // Close after connect
    defer l.Close()
    
    log.Println("Connection Succesfull")
    
    // Authentication
    err = Authentication(l)

    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("Authentication Succesfull")
    
    // Search Request
    result, err := SearchRequest(l)
    
    for _, entry := range result.Entries {
	       fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("userPassword"))
    }
}