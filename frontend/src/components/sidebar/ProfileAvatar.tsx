interface ProfileAvatarProps {
  name: string
}

export function ProfileAvatar({ name }: ProfileAvatarProps) {
  return (
    <button
      aria-label={`Профиль: ${name}`}
      className="flex cursor-not-allowed items-center gap-3 rounded-2xl p-2 text-left"
      disabled
      type="button"
    >
      <span className="grid size-11 shrink-0 place-items-center rounded-full bg-[#00aaff] text-sm font-bold text-white ring-4 ring-white">
        Н
      </span>
      <span className="hidden min-w-0 lg:block">
        <span className="block truncate text-sm font-semibold text-[#1f1f1f]">{name}</span>
        <span className="block text-xs text-[#8a8d91]">Профиль</span>
      </span>
    </button>
  )
}
