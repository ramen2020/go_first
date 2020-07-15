new Vue({
  el: '#app',
  data: {
    tasks: [],
    taskName: '',
    taskMemo: '',
    current: -1,
    options: [{
        value: -1,
        label: 'すべて'
      },
      {
        value: 0,
        label: '問題あり'
      },
      {
        value: 1,
        label: '解決済み'
      }
    ],
    isEntered: false　// false：未入力　true：入力済み
  },
  computed: {
    labels() {
      return this.options.reduce(function (a, b) {
        return Object.assign(a, {
          [b.value]: b.label
        })
      }, {})
    },
    computedTasks() {
      return this.tasks.filter(function (el) {
        var option = this.current < 0 ? true : this.current === el.state
        return option
      }, this)
    },
    validate() {
      var isEnteredTaskName = 0 < this.taskName.length
      this.isEntered = isEnteredTaskName
      return isEnteredTaskName
    }
  },
  // インスタンス作成時の処理
  created: function () {
    this.getAllTasks()
  },
  methods: {
    // 全件取得
    getAllTasks() {
      axios.get('/allTasks')
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            var resultTasks = response.data
            this.tasks = resultTasks
          }
        })
    },
    // 一件取得
    getTaskById(task) {
      axios.get('/task', {
          params: {
            taskID: task.id
          }
        })
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            var resultTasks = response.data
            this.tasks = resultTasks
          }
        })
    },
    // 新規作成
    createTask() {
      const params = new URLSearchParams();
      params.append('taskName', this.taskName)
      params.append('taskMemo', this.taskMemo)

      axios.post('/createTask', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.getAllTasks()
            this.initInputValue()
          }
        })
    },
    // 状態を変更
    changeTaskState(task) {
      const params = new URLSearchParams();
      params.append('taskID', task.id)
      params.append('taskState', task.state)

      axios.post('/changeStateTask', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.getAllTasks()
          }
        })
    },
    // 削除
    deleteTask(task) {
      const params = new URLSearchParams();
      params.append('taskID', task.id)

      axios.post('/deleteTask', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.getAllTasks()
          }
        })
    },
    // 入力値を初期化
    initInputValue() {
      this.current = -1
      this.taskName = ''
      this.taskMemo = ''
    }
  }
})
