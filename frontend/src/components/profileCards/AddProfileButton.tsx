import { Plus } from 'lucide-react'

export function AddProfileButton() {
  return (
    <button
      className="group flex min-h-[282px] cursor-not-allowed flex-col items-center justify-center rounded-[24px] border-2 border-dashed border-[#dfe1e3] bg-transparent p-6 text-center"
      disabled
      type="button"
    >
      <span className="grid size-14 place-items-center rounded-full bg-[#e8f6ff] text-[#00aaff]">
        <Plus aria-hidden="true" className="size-7" strokeWidth={2.2} />
      </span>
      <span className="mt-4 text-sm font-semibold text-[#6f7377]">Добавить профиль</span>
      <span className="mt-1 text-xs text-[#a1a4a7]">Скоро</span>
    </button>
  )
}
