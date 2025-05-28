@echo off
..\..\prometheus-bin\prometheus.exe --config.file=prometheus.yml --web.listen-address=":9090"
pause
