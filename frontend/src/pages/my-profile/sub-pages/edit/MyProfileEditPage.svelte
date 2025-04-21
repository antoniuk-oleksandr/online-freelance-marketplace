<script lang="ts">
  import Loader from '@/common-components/Loader/Loader.svelte'
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import Title from '@/common-components/Title/Title.svelte'
  import MyProfileEditPageForm from './components/MyProfileEditPageForm/MyProfileEditPageForm.svelte'
  import MyProfileEditPageFormContent from './components/MyProfileEditPageFormContent/MyProfileEditPageFormContent.svelte'
  import { fetchProfileDataAndFilters } from './helpers'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import type { FilterParams } from '@/types/FilterParams'

  let request = $state<Promise<[MyProfileEditPageFormData | null, FilterParams | null]>>(
    fetchProfileDataAndFilters(),
  )
</script>

<PaperElement
  styles="flex gap-6 flex-col text-light-palette-text-secondary dark:text-dark-palette-text-secondary h-full flex-grow !ring-0 lg:dark:!ring-1 text-base !bg-transparent dark:!bg-transparent !shadow-none !p-0 md:!p-6 md:!shadow-md dark:!shadow-none md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
  {#await request}
    <Loader />
  {:then [profileData, filterParams]}
    {#if profileData && filterParams}
      <Title text="Edit my profile" />
      <MyProfileEditPageForm data={profileData} filterParams={filterParams}>
        <MyProfileEditPageFormContent filterParams={filterParams}/>
      </MyProfileEditPageForm>
    {/if}
  {/await}
</PaperElement>
