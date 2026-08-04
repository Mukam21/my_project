import { Route, Routes } from 'react-router-dom'

import { PATHS } from '@/config/paths'
import { CartPage } from '@/pages/Cart'
import { HomePage } from '@/pages/Home'

export function AppRouter() {
  const navigations = [
    { path: PATHS.HOME, element: <HomePage /> },
    { path: PATHS.CART, element: <CartPage /> },
  ] as const

  return (
    <Routes>
      {navigations.map(({ path, element }) => (
        <Route key={path} path={path} element={element} />
      ))}
    </Routes>
  )
}
