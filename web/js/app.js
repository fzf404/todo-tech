/*
 * @Author: fzf404
 * @Date: 2022-01-23 18:53:40
 * @LastEditTime: 2022-01-25 15:11:43
 * @Description: 基础
 */

// 路由
$(document).ready(function () {
  // 加载导航栏
  $('#navbar').load('navbar.html', function () {
    // 加载登录框
    $('#login').load('login.html', function () {
      // 全部 Todo 框
      $('#todo').load('todo.html', function () {
        loginVerify() // 验证登录
      })
    })
  })
  // 加载注册框
  $('#regist').load('regist.html')
  // 新增 Todo 框
  $('#add').load('add.html')
  // 更新 Todo 框
  $('#update').load('update.html')
})

// 登录检测
function loginVerify() {
  const token = localStorage.getItem('token')
  if (token == null || token.length == 0) {
    $('#loginButton').click() // 打开登录按钮
    $('.noLogin').removeClass('visually-hidden') // 隐藏 登出按钮 & 头像
    $('.isLogin').addClass('visually-hidden') // 隐藏 登出按钮 & 头像
    $('#todoPage').empty() // 清空全部 Todo
  } else {
    $('.noLogin').addClass('visually-hidden') // 显示 登出按钮 & 头像
    $('.isLogin').removeClass('visually-hidden')
    allTodo() // 获得全部 Todo
  }
}

// 登出
function logOut() {
  localStorage.clear() // 清空 localStorage 以清空 Token
  loginVerify() // 登录验证
}

// 更新 Todo 事件处理
function onUpdate(tid) {
  update_tid = tid // 设置将要编辑的 tid
  // 清空 Tag
  $('#updateModal .tagItem').remove()
  // 获取信息
  data = todo_map.get(tid)
  // 填入信息
  $('#updateTitle').val(data.title)
  $('#updateContent').val(data.content)
  // 插入 Tag
  data.tag.forEach((item) => {
    $('#updateTag').before(`<button class="m-1 btn btn-outline-success tagItem">${item}</button>`)
  })
}
