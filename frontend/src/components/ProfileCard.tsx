import { Heart, MessageCircle, ShoppingBag, Tag } from 'lucide-react'

import type { TestProfile } from '@/types/profile.type'

interface ProfileCardProps {
  profile: TestProfile
}

export function ProfileCard({ profile }: ProfileCardProps) {
  return (
    <article className="group relative min-h-[282px] overflow-hidden rounded-[24px] border border-[#eceeef] bg-white p-5 shadow-[0_2px_16px_rgba(0,0,0,0.04)] transition duration-200 hover:-translate-y-0.5 hover:shadow-[0_10px_28px_rgba(0,0,0,0.08)]">
      <span
        aria-hidden="true"
        className="absolute inset-x-0 top-0 h-1"
        style={{ backgroundColor: profile.accentColor }}
      />

      <header className="flex items-center gap-4">
        <div
          className="relative grid size-14 shrink-0 place-items-center overflow-hidden rounded-full text-sm font-bold text-white"
          style={{ backgroundColor: profile.accentColor }}
        >
          <span>{profile.avatarFallback}</span>
          <img
            alt={`Фото пользователя ${profile.name}`}
            className="absolute inset-0 size-full object-cover"
            src={profile.avatarUrl}
          />
        </div>
        <div className="min-w-0">
          <h2 className="text-base font-bold leading-tight tracking-tight text-[#1f1f1f]">{profile.name}</h2>
          <p className="mt-0.5 text-xs text-[#8a8d91]">{profile.memberSince}</p>
        </div>
      </header>

      <div className="mt-6 grid grid-cols-2 gap-2.5">
        <ProfileStat
          color="#e8f6ff"
          icon={ShoppingBag}
          label="Покупки"
          value={`${profile.purchases} · ${profile.topPurchaseCategory}`}
        />
        <ProfileStat
          color="#f1eafd"
          icon={Tag}
          label="Продажи"
          value={`${profile.sales} · ${profile.saleMessages} отклика`}
        />
        <ProfileStat color="#e7faef" icon={Heart} label="Избранное" value={`${profile.favorites}`} />
        <ProfileStat color="#fff0f2" icon={MessageCircle} label="Чаты" value={`${profile.chats}`} />
      </div>

      <p className="mt-5 text-xs text-[#a1a4a7]">Активность за последний год</p>
    </article>
  )
}

interface ProfileStatProps {
  color: string
  icon: typeof ShoppingBag
  label: string
  value: string
}

function ProfileStat({ color, icon: Icon, label, value }: ProfileStatProps) {
  return (
    <div className="min-w-0 rounded-2xl p-3" style={{ backgroundColor: color }}>
      <div className="flex items-center gap-1.5 text-[#6f7377]">
        <Icon aria-hidden="true" className="size-3.5 shrink-0" strokeWidth={2} />
        <span className="truncate text-[11px] font-medium">{label}</span>
      </div>
      <p className="mt-1.5 text-[11px] font-bold leading-4 text-[#1f1f1f]" title={value}>
        {value}
      </p>
    </div>
  )
}
