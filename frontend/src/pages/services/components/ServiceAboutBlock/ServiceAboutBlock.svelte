<script lang="ts">
  import ServiceAboutBlockLayout from '@/pages/services/components/ServiceAboutBlock/ServiceAboutBlockLayout.svelte'
  import type { Service } from '@/types/Service.ts'
  import Spoiler from '@/common-components/Spoiler/Spoiler.svelte'

  const { description }: Service = $props()
</script>

<ServiceAboutBlockLayout>
  <h3 class="text-xl font-bold text-light-palette-text-primary dark:text-dark-palette-text-primary">
    About
  </h3>
  <Spoiler>
    <div
      class="tiptap text-justify sm:text-pretty text-base h-fit text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
    >
      {@html description}
    </div>
  </Spoiler>
</ServiceAboutBlockLayout>
