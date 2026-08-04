import type { ShoppingBag } from 'lucide-react'

interface ProfileCardStatProps {
  color: string
  icon: typeof ShoppingBag
  label: string
  value: string
}

export function ProfileCardStat({ color, icon: Icon, label, value }: ProfileCardStatProps) {
  return (
    <span className="min-w-0 rounded-2xl p-3" style={{ backgroundColor: color }}>
      <span className="flex items-center gap-1.5 text-[#6f7377]">
        <Icon aria-hidden="true" className="size-3.5 shrink-0" strokeWidth={2} />
        <span className="truncate text-[11px] font-medium">{label}</span>
      </span>
      <span className="mt-1.5 block text-[11px] font-bold leading-4 text-[#1f1f1f]">{value}</span>
    </span>
  )
}
