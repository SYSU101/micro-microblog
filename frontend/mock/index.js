const users = [];

let loggedUserId = -1;
let userIdCounter = 0;

module.exports = {
  'GET /api/sessions': (req, res) => {
    if (loggedUserId > 0) {
      res.status(200).json({ userId: loggedUserId });
    } else {
      res.status(401).json({ errMsg: '用户尚未登录！' });
    }
  },
  'PUT /api/sessions': (req, res) => {
    const { username, password } = req.body;
    let foundUser = null;
    for (const user of users) {
      if (user.username === username) {
        foundUser = user;
        break;
      }
    }
    if (foundUser) {
      loggedUserId = foundUser.id;
      res.status(204).json({ userId: foundUser.id });
    } else {
      res.status(401).json({ errMsg: '用户名或密码不正确' });
    }
  },
  'DELETE /api/sessions': (req, res) => {
    loggedUserId = -1;
    res.status(204).json({});
  },
  'GET /api/users': (req, res) => {
    res.status(200).json({ user: users });
  },
  'POST /api/users': (req, res) => {
    const newUser = req.body;
    for (const user of users) {
      if (user.username === newUser.username) {
        res.status(409).json({ errMsg: '该用户名已存在' });
        return ;
      }
    }
    userIdCounter++;
    users.push({ ...newUser, id: userIdCounter });
    res.status(204).json({})
  },
  'GET /api/user/:id': (req, res) => {
    const { id } = req.params;
    let foundUser = null;
    for (const user of users) {
      if (user.id == id) {
        foundUser = user;
        break;
      }
    }
    if (foundUser) {
      res.status(200).json(foundUser);
    } else {
      res.status(404).json({ errMsg: '找不到对应的用户' });
    }
  },
};