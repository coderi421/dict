import request from '@/utils/request'

export function getDictByKeyword (params) {
  return request({
    url: `/api/v1/dictionary/search`,
    params
  })
}