<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Icon from '@iconify/svelte'
  import { errorStore } from '@/common-stores/error-store'
  import type { ErrorStore } from '@/types/ErrorStore.ts'
  import ErrorComponentLayout from '@/common-components/ErrorComponent/ErrorComponentLayout.svelte'
  import ErrorComponentButton from '@/common-components/ErrorComponent/components/ErrorComponentButton/ErrorComponentButton.svelte'
  import ErrorComponentTitle from '@/common-components/ErrorComponent/components/ErrorComponentTitle/ErrorComponentTitle.svelte'
  import ErrorComponentFirstLineText from '@/common-components/ErrorComponent/components/ErrorComponentFirstLineText/ErrorComponentFirstLineText.svelte'
  import ErrorComponentSecondLineText from '@/common-components/ErrorComponent/components/ErrorComponentSecondLineText/ErrorComponentSecondLineText.svelte'

  let errorData = $state<ErrorStore>()
  errorStore.subscribe((value) => (errorData = value))

  $effect(() => {
    if (errorData && errorData.shown) document.body.style.overflow = 'hidden'
    else document.body.style.overflow = 'auto'
  })
</script>

{#if errorData && errorData.shown}
  <ErrorComponentLayout>
    <PaperElement
      styles="flex text-base !border-0 md:!border !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block flex-col items-center gap-3 p-6 max-w-168"
    >
      <Icon icon="material-symbols:error" class="text-red-500" width="64" height="64" />
      <ErrorComponentTitle error={errorData.error} />
      <ErrorComponentFirstLineText error={errorData.error} />
      <ErrorComponentSecondLineText />
      <ErrorComponentButton />
    </PaperElement>
  </ErrorComponentLayout>
{/if}
