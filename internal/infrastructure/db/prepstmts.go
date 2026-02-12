// db/stmts.go
package db

import (
	"database/sql"
	"log"
	"sync"
)

var (
	stmtMu sync.RWMutex
	stmts  = map[string]*sql.Stmt{}
)

func PrepareCached(query string) (*sql.Stmt, error) {
	stmtMu.RLock()
	s, ok := stmts[query]
	stmtMu.RUnlock()
	if ok {
		return s, nil
	}

	stmtMu.Lock()
	defer stmtMu.Unlock()

	if s, ok = stmts[query]; ok {
		return s, nil
	}

	ps, err := DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	stmts[query] = ps
	log.Println("[DB] Prepared and cached:", query)
	return ps, nil
}

func CloseStmts() {
	stmtMu.Lock()
	defer stmtMu.Unlock()
	for q, s := range stmts {
		_ = s.Close()
		log.Println("[DB] Closed stmt:", q)
	}
}
