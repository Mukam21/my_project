export interface PurchaseRecord {
  title: string
  category: string
  price: number
  date: string
}

export interface TestProfile {
  id: number
  name: string
  joinedAt: string
  avatarUrl: string
  avatarFallback: string
  accentColor: string
  purchases: number
  sales: number
  favoriteCategory: string
  largestPurchase: PurchaseRecord
  smallestPurchase: PurchaseRecord
  annualSpending: number
}
