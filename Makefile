BACKEND_DIR = .
DSN = "user=$(USERNAME) dbname=$(DB) sslmode=disable"
migrate_up:
	cd $(BACKEND_DIR) && goose postgres $(DSN) -dir ./db up 
migrate_down:
	cd $(BACKEND_DIR) && goose postgres $(DSN) -dir ./db down
