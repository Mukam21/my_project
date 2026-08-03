import { createSlice } from '@reduxjs/toolkit'

interface CartState {
  itemCount: number
}

const initialState: CartState = {
  itemCount: 0,
}

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {},
})

export const cartReducer = cartSlice.reducer
