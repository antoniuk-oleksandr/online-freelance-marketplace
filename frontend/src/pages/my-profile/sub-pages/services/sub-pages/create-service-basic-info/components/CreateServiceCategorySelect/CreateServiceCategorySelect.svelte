<script lang="ts">
  import Label from '@/common-components/Label/Label.svelte'
  import NewSelect from '@/common-components/NewSelect/NewSelect.svelte'
  import { getServiceCategoryIdByName } from '@/utils/utils'
  import { getCategorySelectInitialItems } from '../../helpers'
  import CreateServiceCategorySelectLayout from './CreateServiceCategorySelectLayout.svelte'
  import type { Writable } from 'svelte/store'
  import LabeledInputError from '@/common-components/LabeledInput/components/LabeledInputError/LabeledInputError.svelte'
  import type { CreateServiceBasicInfoFormData } from '@/types/CreateServiceBasicInfoFormData'

  type ServiceCreatePageCategorySelectProps = {
    formDataStore: Writable<CreateServiceBasicInfoFormData>
    error?: string | null
  }

  const { formDataStore, error }: ServiceCreatePageCategorySelectProps = $props()

  const initalItems = getCategorySelectInitialItems()

  const updateCategorySelect = (value: string | string[]) => {
    if (Array.isArray(value)) return

    const categoryId = getServiceCategoryIdByName(value)
    formDataStore.update((prev) => ({ ...prev, categoryId }))
  }
</script>

<CreateServiceCategorySelectLayout>
  <Label text="Category" />
  <NewSelect
    menuStyles="!w-full"
    noneSelectedText="Select category"
    {error}
    initialSelectIndex={$formDataStore.categoryId}
    initialSelectValues={initalItems}
    modalHeaderTitle="Select category"
    onSelectChange={(value) => updateCategorySelect(value)}
  />
  <LabeledInputError {error} />
</CreateServiceCategorySelectLayout>
