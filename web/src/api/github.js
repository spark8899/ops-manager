import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url: 'https://api.github.com/repos/spark8899/ops-manager/commits?page=' + page,
    method: 'get'
  })
}

export function Members() {
  return service({
    url: 'https://api.github.com/orgs/spark8899/members',
    method: 'get'
  })
}
