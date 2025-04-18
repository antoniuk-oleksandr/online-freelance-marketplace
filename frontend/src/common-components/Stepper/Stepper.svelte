<script lang="ts">
  import Icon from '@iconify/svelte'
  import StepperLayout from '@/common-components/Stepper/StepperLayout.svelte'
  import StepperElement from '@/common-components/Stepper/components/StepperElement/StepperElement.svelte'
  import type { StepperItem } from '@/types/StepperItem.ts'

  type StepperProps = {
    steps: StepperItem[]
    activeStepIndex: number
    styles?: string
  }

  const { steps, activeStepIndex, styles }: StepperProps = $props()
</script>

<StepperLayout {styles}>
  {#each steps as step, index}
    <StepperElement {step} {index} {activeStepIndex} />
    {#if index !== steps.length - 1}
      <Icon width={18} height={18} icon="hugeicons:arrow-right-01" />
    {/if}
  {/each}
</StepperLayout>
