const signupForm = document.getElementById('signup')
const loginForm = document.getElementById('login')

function show(node, message, ok=true){
  node.textContent = message
  node.className = 'msg ' + (ok ? 'ok' : 'err')
}

signupForm.addEventListener('submit', async (e) => {
  e.preventDefault()
  const f = e.target
  const data = {
    username: f.username.value,
    email: f.email.value,
    password: f.password.value,
  }
  const msg = document.getElementById('signup-msg')
  try{
    const res = await fetch('/api/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })
    if(res.ok){
      show(msg, 'Account created successfully', true)
      f.reset()
    } else {
      const text = await res.text()
      show(msg, 'Failed to create account: ' + text, false)
    }
  }catch(err){
    show(msg, 'Network error: ' + err.message, false)
  }
})

loginForm.addEventListener('submit', async (e) => {
  e.preventDefault()
  const f = e.target
  const data = {
    email: f.email.value,
    password: f.password.value,
  }
  const msg = document.getElementById('login-msg')
  try{
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })
    if(res.status === 401){
      show(msg, 'Username or password is incorrect', false)
      return
    }
    if(res.ok){
      const body = await res.json()
      show(msg, 'Login successful', true)
      f.reset()
    } else {
      const text = await res.text()
      show(msg, 'Login error: ' + text, false)
    }
  }catch(err){
    show(msg, 'Network error: ' + err.message, false)
  }
})
