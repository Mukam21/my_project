import type { PurchaseRecord } from '@/types/profile.type'

export function getAnnualSpending(purchases: PurchaseRecord[]) {
  return purchases.reduce((total, purchase) => total + purchase.price, 0)
}

export function getLargestPurchase(purchases: PurchaseRecord[]) {
  return purchases.reduce((largest, purchase) => (purchase.price > largest.price ? purchase : largest))
}

export function getSmallestPurchase(purchases: PurchaseRecord[]) {
  return purchases.reduce((smallest, purchase) => (purchase.price < smallest.price ? purchase : smallest))
}
