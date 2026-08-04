import type { ReactNode } from 'react'
import { ArrowDownRight, ArrowUpRight, CalendarDays, Sparkles } from 'lucide-react'

import { PurchaseDetailsCard } from '@/components/PurchaseDetailsCard'
import { ProfileImage } from '@/components/ProfileImage'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import type { TestProfile } from '@/types/profile.type'
import {formatDate} from '@/utils/formatterDate'
import {formatCurrency} from '@/utils/formatterNumber'
interface ProfileDetailsDialogProps {
  children: ReactNode
  profile: TestProfile
}

export function ProfileDetailsDialog({ children, profile }: ProfileDetailsDialogProps) {
  return (
    <Dialog>
      {children}
      <DialogContent className="max-h-[calc(100dvh-2rem)] overflow-y-auto rounded-[28px] p-5 sm:max-w-2xl sm:p-7">
        <DialogHeader className="pr-10">
          <div className="flex items-center gap-4">
            <ProfileImage profile={profile} size="large" />
            <div>
              <DialogTitle className="text-xl font-bold text-[#1f1f1f] sm:text-2xl">{profile.name}</DialogTitle>
              <DialogDescription className="mt-1 flex items-center gap-1.5 text-[#8a8d91]">
                <CalendarDays aria-hidden="true" className="size-3.5" />
                С {formatDate(profile.joinedAt)}
              </DialogDescription>
            </div>
          </div>
        </DialogHeader>

        <div className="grid grid-cols-3 gap-2.5">
          <ModalStat label="Покупок" value={`${profile.purchases}`} />
          <ModalStat label="Продаж" value={`${profile.sales}`} />
          <ModalStat label="Потрачено" value={formatCurrency(profile.annualSpending)} />
        </div>

        <section className="rounded-2xl bg-[#e8f6ff] p-4">
          <div className="flex items-center gap-2 text-xs font-medium text-[#6f7377]">
            <Sparkles aria-hidden="true" className="size-4 text-[#00aaff]" />
            Любимая категория покупок
          </div>
          <p className="mt-2 text-base font-bold text-[#1f1f1f]">{profile.favoriteCategory}</p>
        </section>

        <div className="grid gap-3 sm:grid-cols-2">
          <PurchaseDetailsCard
            backgroundColor="bg-[#f1eafd]"
            icon={ArrowUpRight}
            iconColor="text-[#965eeb]"
            label="Самая большая покупка"
            purchase={profile.largestPurchase}
          />
          <PurchaseDetailsCard
            backgroundColor="bg-[#e7faef]"
            icon={ArrowDownRight}
            iconColor="text-[#00b956]"
            label="Самая маленькая покупка"
            purchase={profile.smallestPurchase}
          />
        </div>
      </DialogContent>
    </Dialog>
  )
}

function ModalStat({ label, value }: { label: string; value: string }) {
  return (
    <div className="min-w-0 rounded-2xl bg-[#f5f5f5] p-3 sm:p-4">
      <p className="text-[11px] text-[#8a8d91] sm:text-xs">{label}</p>
      <p className="mt-1 truncate text-sm font-bold text-[#1f1f1f] sm:text-base" title={value}>
        {value}
      </p>
    </div>
  )
}
