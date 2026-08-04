import { CircleDollarSign, ShoppingBag, Sparkles, Tag } from 'lucide-react'

import { DialogTrigger } from '@/components/ui/dialog'
import { ProfileCardStat } from './ProfileCardStat'
import { ProfileDetailsDialog } from './ProfileDetailsDialog'
import { ProfileImage } from './ProfileImage'
import type { TestProfile } from '@/types/profile.type'
import {formatDateSince} from '@/utils/formatterDate'
import {formatCurrency} from '@/utils/formatterNumber'
import { getAnnualSpending } from '@/utils/profileAnalytics'


interface ProfileCardProps {
  profile: TestProfile
}

export function ProfileCard({ profile }: ProfileCardProps) {
  return (
    <ProfileDetailsDialog profile={profile}>
      <DialogTrigger className="cursor-pointer group relative min-h-[282px] overflow-hidden rounded-[24px] border border-[#eceeef] bg-white p-5 text-left shadow-[0_2px_16px_rgba(0,0,0,0.04)] transition duration-200 hover:-translate-y-0.5 hover:shadow-[0_10px_28px_rgba(0,0,0,0.08)] focus-visible:outline-3 focus-visible:outline-offset-3 focus-visible:outline-[#00aaff]">
        <span
          aria-hidden="true"
          className="absolute inset-x-0 top-0 h-1"
          style={{ backgroundColor: profile.accentColor }}
        />

        <div className="flex items-center gap-4">
          <ProfileImage profile={profile} />
          <div className="min-w-0">
            <h2 className="text-base font-bold leading-tight tracking-tight text-[#1f1f1f]">{profile.name}</h2>
            <p className="mt-0.5 text-xs text-[#8a8d91]">{formatDateSince(profile.joinedAt)}</p>
          </div>
        </div>

        <div className="mt-6 grid grid-cols-2 gap-2.5">
          <ProfileCardStat color="#e8f6ff" icon={ShoppingBag} label="Покупки" value={`${profile.purchases.length}`} />
          <ProfileCardStat color="#f1eafd" icon={Tag} label="Продажи" value={`${profile.sales.length}`} />
          <ProfileCardStat
            color="#e7faef"
            icon={Sparkles}
            label="Любимая категория"
            value={profile.favoriteCategory}
          />
          <ProfileCardStat
            color="#fff0f2"
            icon={CircleDollarSign}
            label="Потрачено"
            value={formatCurrency(getAnnualSpending(profile.purchases))}
          />
        </div>
      </DialogTrigger>
    </ProfileDetailsDialog>
  )
}
