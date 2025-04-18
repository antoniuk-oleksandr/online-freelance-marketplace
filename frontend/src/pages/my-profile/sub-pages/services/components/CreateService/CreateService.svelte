<script lang="ts">
  import CreateServiceLayout from './CreateServiceLayout.svelte'
  import type { LayoutProps } from '@/types/LayoutProps'
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Stepper from '@/common-components/Stepper/Stepper.svelte'
  import { makeCreateServiceSteps, makeCreateServiceTitleData } from './helpers'
  import CreateServiceForm from '../CreateServiceForm/CreateServiceForm.svelte'
  import { z } from 'zod'
  import Title from '@/common-components/Title/Title.svelte'
  import BackButton from '@/common-components/BackButton/BackButton.svelte'

  type CreateServiceProps = LayoutProps & {
    activeStepIndex: number
    onSubmit: (data: any) => void
    schema: z.ZodObject<any>
  }

  const { children, activeStepIndex, ...rest }: CreateServiceProps = $props()
  const steps = makeCreateServiceSteps(activeStepIndex)
  const { title, subtitle } = makeCreateServiceTitleData(activeStepIndex)
</script>

<CreateServiceLayout>
  <Stepper styles="justify-center" {steps} {activeStepIndex} />
  <PaperElement
    styles="flex text-light-palette-text-secondary dark:text-dark-palette-text-secondary h-full flex-grow !ring-0 lg:dark:!ring-1 text-base !bg-transparent dark:!bg-transparent !shadow-none !p-0 md:!p-6 md:!shadow-md dark:!shadow-none md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
  >
    <CreateServiceForm {...rest}>
      <Title text={title} description={subtitle} capitalize={false} />
      {@render children()}
      {#if activeStepIndex > 0}
        <BackButton styles="w-fit -mt-3 self-end" />
      {/if}
    </CreateServiceForm>
  </PaperElement>
</CreateServiceLayout>
