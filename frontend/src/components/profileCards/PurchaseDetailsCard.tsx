import type { LucideIcon } from 'lucide-react'

import type { PurchaseRecord } from '@/types/profile.type'
import {formatDate} from '@/utils/formatterDate'
import {formatCurrency} from '@/utils/formatterNumber'
interface PurchaseDetailsCardProps {
  backgroundColor: string
  icon: LucideIcon
  iconColor: string
  label: string
  purchase: PurchaseRecord
}

export function PurchaseDetailsCard({
  backgroundColor,
  icon: Icon,
  iconColor,
  label,
  purchase,
}: PurchaseDetailsCardProps) {
  return (
    <section className={`rounded-2xl p-4 ${backgroundColor}`}>
      <div className="flex items-center gap-2 text-xs font-medium text-[#6f7377]">
        <Icon aria-hidden="true" className={`size-4 ${iconColor}`} />
        {label}
      </div>
      <p className="mt-3 text-base font-bold text-[#1f1f1f]">{purchase.title}</p>
      <p className="mt-0.5 text-xs text-[#6f7377]">{purchase.category}</p>
      <div className="mt-4 flex items-end justify-between gap-3">
        <p className="text-lg font-bold text-[#1f1f1f]">{formatCurrency(purchase.price)}</p>
        <p className="text-right text-[11px] text-[#8a8d91]">{formatDate(purchase.date)}</p>
      </div>
    </section>
  )
}
