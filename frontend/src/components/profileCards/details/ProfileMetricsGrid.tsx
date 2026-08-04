import { Heart, MapPin, MessageCircle, PackagePlus, Star, ThumbsUp } from 'lucide-react'

import type { TestProfile } from '@/types/profile.type'

interface ProfileMetricsGridProps {
  profile: TestProfile
}

export function ProfileMetricsGrid({ profile }: ProfileMetricsGridProps) {
  const { chatsCount, metrics } = profile
  const items = [
    { icon: Heart, label: 'Лайки профиля', value: metrics.likes },
    { icon: Star, label: 'Отзывы', value: metrics.reviews },
    { icon: ThumbsUp, label: 'Избранные', value: metrics.favoriteListings },
    { icon: PackagePlus, label: 'Созданные объявления', value: metrics.createdListings },
    { icon: MessageCircle, label: 'Чаты', value: chatsCount },
    { icon: MapPin, label: 'Город', value: metrics.city },
  ]

  return (
    <section>
      <h3 className="text-base font-bold text-[#1f1f1f]">Профиль</h3>
      <div className="mt-3 grid grid-cols-2 gap-2.5 sm:grid-cols-3">
        {items.map(({ icon: Icon, label, value }) => (
          <div key={label} className="min-w-0 rounded-2xl bg-[#f5f5f5] p-3">
            <div className="flex items-center gap-1.5 text-[#8a8d91]">
              <Icon aria-hidden="true" className="size-3.5" />
              <span className="truncate text-[11px]">{label}</span>
            </div>
            <p className="mt-1.5 truncate text-sm font-bold text-[#1f1f1f]" title={String(value)}>
              {value}
            </p>
          </div>
        ))}
      </div>
      <p className="mt-3 text-xs text-[#8a8d91]">
        Рейтинг {metrics.rating.toFixed(1)} · Активных дней за год: {metrics.activeDays}
      </p>
    </section>
  )
}
