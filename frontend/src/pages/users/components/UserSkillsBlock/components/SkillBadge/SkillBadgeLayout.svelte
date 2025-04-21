<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { Link } from 'svelte-routing'
  import { getHoverClass } from '@/utils/utils'

  type SkillBadgeLayoutProps = {
    id: number
  } & LayoutProps

  let { id, children }: SkillBadgeLayoutProps = $props()
  const hoverClass = getHoverClass()
</script>

<Link
  to={'/search?skill=' + id}
  class={`rounded-xl grid place-items-center text-center border border-light-palette-divider dark:border-dark-palette-divider px-3 py-1 bg-light-palette-background-block dark:bg-dark-palette-background-block ${hoverClass}}`}
>
  {@render children()}
</Link>
