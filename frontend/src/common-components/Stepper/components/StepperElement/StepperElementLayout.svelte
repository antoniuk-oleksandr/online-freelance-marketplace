<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { Link } from 'svelte-routing'

  type StepperElementProps = LayoutProps & {
    active: boolean
    link?: string
  }

  const { children, active, link }: StepperElementProps = $props()
</script>

{#snippet data()}
  <div class="flex items-center gap-1 {active ? 'text-cyan-500' : ''}">
    {@render children()}
  </div>
{/snippet}

{#if link !== undefined}
  <Link to={link}>
    {@render data()}
  </Link>
{:else}
  {@render data()}
{/if}
