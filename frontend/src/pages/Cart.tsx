import { Link } from 'react-router-dom'

import { PATHS } from '@/config/paths'

export function CartPage() {
  return (
    <main className="mx-auto flex min-h-screen max-w-3xl flex-col justify-center px-6">
      <h1 className="text-4xl font-bold text-neutral-900">Корзина</h1>
      <Link className="mt-4 w-fit text-[#00aaef] hover:underline" to={PATHS.HOME}>
        На главную
      </Link>
    </main>
  )
}
