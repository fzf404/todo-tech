/*
 * @Author: fzf404
 * @Date: 2022-01-23 20:58:38
 * @LastEditTime: 2022-01-25 15:10:22
 * @Description: ajax 请求
 */

// const BASE_URL = 'http://127.0.0.1:8080'
const BASE_URL = '/api/v2'
// 数据
let todo_map = new Map() // 存储 TODO 信息
let update_tid = 0 // 修改操作时记录 tid 值
// 分页
let page = 1

/**
 * @description: 登录
 */
function login() {
  $.ajax({
    url: `${BASE_URL}/user/login`,
    type: 'POST',
    contentType: 'application/json',
    data: JSON.stringify({
      username: $('#loginUsername').val(),
      password: $('#loginPassword').val(),
    }),
    success: function (res) {
      alert(res.msg) // 响应信息
      if (res.code == 200) {
        localStorage.setItem('token', res.data.token) // 存储 Token
        $('#closeLogin').click() // 隐藏登录框
        $('.noLogin').addClass('visually-hidden') // 显示 登出按钮 & 头像
        $('.isLogin').removeClass('visually-hidden')
        allTodo() // 获得全部 Todo
      }
    },
  })
}

/**
 * @description: 注册
 */
function regist() {
  $.ajax({
    url: `${BASE_URL}/user/regist`,
    type: 'POST',
    contentType: 'application/json',
    data: JSON.stringify({
      username: $('#registUsername').val(),
      email: $('#registEmail').val(),
      password: $('#registPassword').val(),
    }),
    success: function (res) {
      alert(res.msg) // 响应信息
      if (res.code == 200) {
        $('#loginButton').click() // 显示登录框
      }
    },
  })
}

/**
 * @description: 全部 Todo
 * @param {*} page
 */
function allTodo() {
  $('#todoPage').empty() // 清空 Todo
  $('#searchBack').addClass('visually-hidden') // 隐藏返回按钮
  for (let p = 1; p <= page; p++) {
    $.ajax({
      url: `${BASE_URL}/todo/all?page=${p}`,
      type: 'GET',
      // 设置请求头
      beforeSend: (req) => {
        token = window.localStorage.getItem('token')
        req.setRequestHeader('x-token', token)
      },
      success: function (res) {
        if (res.code == 200) {
          res.data.forEach((item) => {
            todo_map.set(item.tid, item) // 储存信息
            // 处理标签
            let badges = ''
            item.tag.forEach((i) => {
              badges += `<span class="badge bg-success me-1">${i}</span>`
            })
            // 添加进页面
            $('#todoPage').append(`
            <div class="card m-2" style="max-width: 260px">
              <div class="card-body">
                <h5 class="card-title">${item.title}</h5>
                <span class="card-text">${item.content}</span>
                <br />
                ${badges}
              </div>
              <div class="card-footer">
                <button class="btn btn-sm btn-outline-primary" data-bs-toggle="modal" data-bs-target="#updateModal" onclick="onUpdate(${item.tid})">编辑</button>
                <button class="btn btn-sm btn-outline-danger" onclick="removeTodo(${item.tid})">删除</button>
              </div>
            </div>
          `)
          })
          if (res.data.length == 8) {
            $('#moreTodo').removeClass('visually-hidden') // 显示更多按钮
          }
        } else {
          alert(res.msg) // 响应信息
        }
      },
    })
  }
}

/**
 * @description: 搜索 Todo
 */
function searchTodo() {
  $.ajax({
    url: `${BASE_URL}/todo/search?title=${$('#searchTodo').val()}`,
    type: 'GET',
    // 设置请求头
    beforeSend: (req) => {
      token = window.localStorage.getItem('token')
      req.setRequestHeader('x-token', token)
    },
    success: function (res) {
      if (res.code == 200) {
        $('#todoPage').empty() // 清空 Todo
        $('#moreTodo').addClass('visually-hidden') // 隐藏更多按钮
        $('#searchBack').removeClass('visually-hidden') // 显示返回按钮

        res.data.forEach((item) => {
          todo_map.set(item.tid, item) // 储存信息
          // 处理标签
          let badges = ''
          item.tag.forEach((i) => {
            badges += `<span class="badge bg-success me-1">${i}</span>`
          })
          // 添加进页面
          $('#todoPage').append(`
          <div class="card m-2" style="max-width: 260px">
            <div class="card-body">
              <h5 class="card-title">${item.title}</h5>
              <span class="card-text">${item.content}</span>
              <br />
              ${badges}
            </div>
            <div class="card-footer">
              <button class="btn btn-sm btn-outline-primary" data-bs-toggle="modal" data-bs-target="#updateModal" onclick="onUpdate(${item.tid})">编辑</button>
              <button class="btn btn-sm btn-outline-danger" onclick="removeTodo(${item.tid})">删除</button>
            </div>
          </div>
        `)
        })
      } else {
        alert(res.msg) // 响应信息
      }
    },
  })
}
/**
 * @description: 获取更多 Todo
 */
function moreTodo() {
  page += 1
  $.ajax({
    url: `${BASE_URL}/todo/all?page=${page}`,
    type: 'GET',
    // 设置请求头
    beforeSend: (req) => {
      token = window.localStorage.getItem('token')
      req.setRequestHeader('x-token', token)
    },
    success: function (res) {
      if (res.code == 200) {
        res.data.forEach((item) => {
          todo_map.set(item.tid, item) // 储存信息
          // 处理标签
          let badges = ''
          item.tag.forEach((i) => {
            badges += `<span class="badge bg-success me-1">${i}</span>`
          })
          // 添加进页面
          $('#todoPage').append(`
          <div class="card m-2" style="max-width: 260px">
            <div class="card-body">
              <h5 class="card-title">${item.title}</h5>
              <span class="card-text">${item.content}</span>
              <br />
              ${badges}
            </div>
            <div class="card-footer">
              <button class="btn btn-sm btn-outline-primary" data-bs-toggle="modal" data-bs-target="#updateModal" onclick="onUpdate(${item.tid})">编辑</button>
              <button class="btn btn-sm btn-outline-danger" onclick="removeTodo(${item.tid})">删除</button>
            </div>
          </div>
        `)
        })
        // 判断是否加载更多
        if (res.data.length == 8) {
          $('#moreTodo').removeClass('visually-hidden') // 显示更多按钮
        } else {
          $('#moreTodo').addClass('visually-hidden') // 隐藏按钮
        }
      } else {
        alert(res.msg) // 响应信息
        $('#moreTodo').addClass('visually-hidden') // 隐藏按钮
        page -= 1
      }
    },
  })
}

/**
 * @description: 增加 Todo
 */
function addTodo() {
  // 批量读取 tag
  let tags = []
  $('#addModal .tagItem').each((i, e) => {
    tags.push($(e).text())
  })
  $.ajax({
    url: `${BASE_URL}/todo/add`,
    type: 'POST',
    contentType: 'application/json',
    // 设置请求头
    beforeSend: (req) => {
      token = window.localStorage.getItem('token')
      req.setRequestHeader('x-token', token)
    },
    data: JSON.stringify({
      title: $('#addTitle').val(),
      content: $('#addContent').val(),
      tag: tags,
    }),
    success: function (res) {
      alert(res.msg) // 响应信息
      if (res.code == 200) {
        $('#closeAdd').click() // 关闭模态框
        allTodo() // 重新获得全部 Todo
        // 清空输入框
        $('#addTitle').val('')
        $('#addContent').val('')
        $('#addModal .tagItem').remove()
      }
    },
  })
}

/**
 * @description: 更新 Todo
 */
function updateTodo() {
  // 批量读取 tag
  let tags = []
  $('#updateModal .tagItem').each((i, e) => {
    tags.push($(e).text())
  })
  $.ajax({
    url: `${BASE_URL}/todo/update`,
    type: 'POST',
    contentType: 'application/json',
    // 设置请求头
    beforeSend: (req) => {
      token = window.localStorage.getItem('token')
      req.setRequestHeader('x-token', token)
    },
    data: JSON.stringify({
      tid: update_tid,
      title: $('#updateTitle').val(),
      content: $('#updateContent').val(),
      tag: tags,
    }),
    success: function (res) {
      alert(res.msg) // 响应信息
      if (res.code == 200) {
        $('#closeUpdate').click() // 关闭模态框
        allTodo() // 重新获得全部 Todo
      }
    },
  })
}

/**
 * @description: 删除 Todo
 * @param {*} tid
 */
function removeTodo(tid) {
  $.ajax({
    url: `${BASE_URL}/todo/remove`,
    type: 'POST',
    contentType: 'application/json',
    // 设置请求头
    beforeSend: (req) => {
      token = window.localStorage.getItem('token')
      req.setRequestHeader('x-token', token)
    },
    data: JSON.stringify({
      tid: tid,
    }),
    success: function (res) {
      alert(res.msg) // 响应信息
      if (res.code == 200) {
        allTodo() // 重新获得全部 Todo
      }
    },
  })
}
