<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'
  import { dndzone } from 'svelte-dnd-action'

  type CreateServiceQuestionListLayoutProps = LayoutProps & {
    items: { text: string; id: number }[]
    setQuestions: (value: { text: string; id: number }[]) => void
  }

  const { children, items, setQuestions }: CreateServiceQuestionListLayoutProps = $props()

  const handleConsider = (evt: any) => setQuestions(evt.detail.items)
  const handleFinalize = (evt: any) => setQuestions(evt.detail.items)
</script>

<section
  class="flex flex-col gap-6 !outline-none"
  use:dndzone={{ items, flipDurationMs: 200 }}
  onconsider={handleConsider}
  onfinalize={handleFinalize}
>
  {@render children()}
</section>
