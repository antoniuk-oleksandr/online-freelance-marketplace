<script lang="ts">
  import LabeledInput from '@/common-components/LabeledInput/LabeledInput.svelte'
  import MyProfileEditPageInputsBlockLayout from './MyProfileEditPageInputsBlockLayout.svelte'
  import MyProfileEditPageAboutInput from '../MyProfileEditPageAboutInput/MyProfileEditPageAboutInput.svelte'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import type { Writable } from 'svelte/store'
  import type { FilterParams } from '@/types/FilterParams'
  import MyProfileEditPageLanguagesSkillsBlock from '../MyProfileEditPageLanguagesSkillsBlock/MyProfileEditPageLanguagesSkillsBlock.svelte'
  import { getContext } from 'svelte'
  import { getMyProfileEditPageInputsData } from '../../helpers'
  import type { MyProfileEditPageInputData } from '@/types/MyProfileEditPageInputData'

  type MyProfileEditPageInputsBlockProps = {
    formDataStore: Writable<MyProfileEditPageFormData>
    filterParams: FilterParams
  }

  const formErrorsStore = getContext<Writable<Record<string, string[]>>>('formErrorsStore')
  let inputsData = $state<MyProfileEditPageInputData[] | undefined>()
  formErrorsStore.subscribe((value) => (inputsData = getMyProfileEditPageInputsData(value)))

  const props: MyProfileEditPageInputsBlockProps = $props()
</script>

<MyProfileEditPageInputsBlockLayout>
  {#if inputsData !== undefined}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      {#each inputsData as item}
        <LabeledInput {...item} />
      {/each}
    </div>
    <MyProfileEditPageAboutInput {...props} />

    <MyProfileEditPageLanguagesSkillsBlock {formErrorsStore} {...props} type="languages" />
    <MyProfileEditPageLanguagesSkillsBlock {formErrorsStore} {...props} type="skills" />
  {/if}
</MyProfileEditPageInputsBlockLayout>
