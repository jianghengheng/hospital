@echo off
echo 正在启动服务...

:: 创建 tmp 目录（如果不存在）
if not exist "tmp" mkdir tmp

:: 运行 air
air 