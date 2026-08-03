import { Heart, Settings } from 'lucide-react'

import { ProfileAvatar } from '@/components/ProfileAvatar'

const sidebarItems = [
  { label: 'Избранное', icon: Heart },
  { label: 'Настройки', icon: Settings },
] as const

export function Sidebar() {
  return (
    <aside className="sticky top-0 flex h-dvh w-[82px] shrink-0 flex-col border-r border-[#e6e7e8] bg-[#f5f5f5] px-3 py-6 lg:w-[236px] lg:px-5">
      <div className="flex items-center gap-3 px-2">
        <span aria-hidden="true" className="grid grid-cols-2 gap-[2px]">
          <span className="size-[7px] rounded-full bg-[#00aaff]" />
          <span className="size-[7px] rounded-full bg-[#965eeb]" />
          <span className="size-[7px] rounded-full bg-[#04e061]" />
          <span className="size-[7px] rounded-full bg-[#ff4053]" />
        </span>
        <span className="hidden text-base font-bold tracking-tight text-[#1f1f1f] lg:block">Итоги года</span>
      </div>

      <nav aria-label="Боковая навигация" className="mt-12 space-y-2">
        {sidebarItems.map(({ label, icon: Icon }) => (
          <button
            key={label}
            className="flex w-full cursor-not-allowed items-center justify-center gap-3 rounded-xl px-3 py-3 text-[#6f7377] lg:justify-start"
            disabled
            type="button"
          >
            <Icon aria-hidden="true" className="size-5" strokeWidth={1.8} />
            <span className="hidden text-sm font-medium lg:block">{label}</span>
          </button>
        ))}
      </nav>

      <div className="mt-auto">
        <ProfileAvatar name="Никита" />
      </div>
    </aside>
  )
}
