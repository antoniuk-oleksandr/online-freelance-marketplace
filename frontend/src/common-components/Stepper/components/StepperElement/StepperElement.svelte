<script lang="ts">
  import Icon from '@iconify/svelte'
  import StepperElementLayout from '@/common-components/Stepper/components/StepperElement/StepperElementLayout.svelte'

  type StepperElementProps = {
    step: { text: string; link?: string }
    index: number
    activeStepIndex: number
  }

  const { step, index, activeStepIndex }: StepperElementProps = $props()
</script>

<StepperElementLayout link={step.link} active={index === activeStepIndex}>
  <Icon width={24} height={24} icon="mdi:number-{index + 1}-circle" />
  <span>{step.text}</span>
</StepperElementLayout>
