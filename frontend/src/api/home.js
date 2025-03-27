import request from '@/utils/request'

export function getDictByKeyword (params) {
  return request({
    url: `/api/v1/dictionary/search`,
    params
  })
}

export function getAllCategory () {
  return request({
    url: `/api/v1/category/all`
  })
}