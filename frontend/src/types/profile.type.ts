export interface PurchaseRecord {
  title: string
  category: string
  price: number
  date: string
}

export interface SaleRecord extends PurchaseRecord {
  inquiriesCount: number
}

export interface ListingViewRecord {
  title: string
  category: string
  likes: number
  viewedAt: string
  viewCount: number
}

export interface ProfileMetrics {
  activeDays: number
  city: string
  createdListings: number
  favoriteListings: number
  likes: number
  rating: number
  reviews: number
}

export interface TestProfile {
  id: number
  name: string
  joinedAt: string
  avatarUrl: string
  avatarFallback: string
  accentColor: string
  chatsCount: number
  favoriteCategory: string
  metrics: ProfileMetrics
  purchases: PurchaseRecord[]
  sales: SaleRecord[]
  listingViews: ListingViewRecord[]
}
