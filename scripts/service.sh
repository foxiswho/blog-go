#!/bin/bash
# 当前文件目录
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
# 项目启动/停止/重启管理脚本
# 使用方式: ./service.sh [start|stop|restart|status]

# ======================== 配置区域 (请根据你的项目修改) ========================
# 项目名称 (用于标识进程)
APP_NAME="blogGo"
# 编译后的可执行文件路径 (绝对路径)
APP_BIN=$SCRIPT_DIR
# 日志文件路径
LOG_FILE="${SCRIPT_DIR}/blogGo.log"
# PID 文件路径 (记录进程ID)
PID_FILE="${SCRIPT_DIR}/${APP_NAME}.pid"
# 启动参数 (如有需要)
APP_ARGS=""
# ======================== 配置区域结束 ========================

# 检查可执行文件是否存在
check_app_exist() {
    if [ ! -f "${APP_BIN}" ]; then
        echo "错误: 可执行文件 ${APP_BIN} 不存在!"
        exit 1
    fi
}

# 检查进程是否运行
check_running() {
    if [ -f "${PID_FILE}" ]; then
        PID=$(cat "${PID_FILE}")
        if ps -p "${PID}" > /dev/null 2>&1; then
            return 0  # 运行中
        else
            rm -f "${PID_FILE}"  # PID文件存在但进程已退出，清理文件
        fi
    fi
    return 1  # 未运行
}

# 启动服务
start() {
    if check_running; then
        echo "${APP_NAME} 已在运行 (PID: $(cat ${PID_FILE}))"
        return 0
    fi

    check_app_exist
    echo "正在启动 ${APP_NAME}..."

    # 后台运行程序，输出重定向到日志文件
    nohup "${APP_BIN}" ${APP_ARGS} > "${LOG_FILE}" 2>&1 &
    echo $! > "${PID_FILE}"

    # 短暂延迟后检查是否启动成功
    sleep 1
    if check_running; then
        echo "${APP_NAME} 启动成功 (PID: $(cat ${PID_FILE}))"
    else
        echo "错误: ${APP_NAME} 启动失败，请查看日志 ${LOG_FILE}"
        rm -f "${PID_FILE}"
        exit 1
    fi
}

# 停止服务
stop() {
    if ! check_running; then
        echo "${APP_NAME} 未运行"
        return 0
    fi

    PID=$(cat "${PID_FILE}")
    echo "正在停止 ${APP_NAME} (PID: ${PID})..."

    # 优雅停止 (先尝试 SIGTERM，失败则 SIGKILL)
    kill "${PID}" > /dev/null 2>&1
    sleep 2

    if ps -p "${PID}" > /dev/null 2>&1; then
        echo "优雅停止失败，强制终止 ${APP_NAME}..."
        kill -9 "${PID}" > /dev/null 2>&1
    fi

    rm -f "${PID_FILE}"
    echo "${APP_NAME} 已停止"
}

# 查看状态
status() {
    if check_running; then
        echo "${APP_NAME} 正在运行 (PID: $(cat ${PID_FILE}))"
    else
        echo "${APP_NAME} 未运行"
    fi
}

# 重启服务
restart() {
    stop
    start
}

# 显示帮助信息
usage() {
    echo "使用方法: $0 [start|stop|restart|status]"
    echo "  start   - 启动服务"
    echo "  stop    - 停止服务"
    echo "  restart - 重启服务"
    echo "  status  - 查看服务状态"
    exit 1
}

# 主逻辑
case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        restart
        ;;
    status)
        status
        ;;
    *)
        usage
        ;;
esac

exit 0