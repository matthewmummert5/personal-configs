#!/bin/bash

# Load configuration
source $HOME/restic-env.sh

# Set log file
LOGFILE="$HOME/.restic-backup.log"

# Function to log messages
log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') - $1" | tee -a "$LOGFILE"
}

# Start backup
log "Starting backup process"

# Perform backup with exclusions
restic backup \
    --verbose \
    --files-from /home/username/.backup.list \
    --tag "$(date +%Y-%m-%d)" \
    2>&1 | tee -a "$LOGFILE"

if [ $? -eq 0 ]; then
    log "Backup completed successfully"
else
    log "Backup failed with exit code $?"
    exit 1
fi

# Cleanup: keep last 7 daily, 4 weekly, 6 monthly snapshots
log "Starting cleanup process"
restic forget \
    --keep-daily 7 \
    --keep-weekly 4 \
    --keep-monthly 6 \
    --prune \
    2>&1 | tee -a "$LOGFILE"

log "Backup and cleanup completed"
