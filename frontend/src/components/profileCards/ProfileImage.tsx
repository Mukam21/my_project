import type { TestProfile } from '@/types/profile.type'

interface ProfileImageProps {
  profile: TestProfile
  size?: 'default' | 'large'
}

export function ProfileImage({ profile, size = 'default' }: ProfileImageProps) {
  const sizeClass = size === 'large' ? 'size-16' : 'size-14'

  return (
    <span
      className={`relative grid ${sizeClass} shrink-0 place-items-center overflow-hidden rounded-full text-sm font-bold text-white`}
      style={{ backgroundColor: profile.accentColor }}
    >
      <span>{profile.avatarFallback}</span>
      <img
        alt={`Фото пользователя ${profile.name}`}
        className="absolute inset-0 size-full object-cover"
        src={profile.avatarUrl}
      />
    </span>
  )
}
