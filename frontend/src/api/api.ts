import axios from 'axios'

const DEFAULT_BASE_URL = 'http://localhost:8080/api'

export const API_BASE_URL = import.meta.env.VITE_API_URL?.trim() || DEFAULT_BASE_URL

export const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})
