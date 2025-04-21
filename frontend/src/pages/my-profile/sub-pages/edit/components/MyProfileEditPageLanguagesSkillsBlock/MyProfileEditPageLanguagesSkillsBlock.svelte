<script lang="ts">
  import Label from '@/common-components/Label/Label.svelte'
  import type { FilterParams } from '@/types/FilterParams'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import type { Writable } from 'svelte/store'
  import MyProfileEditPageLanguagesSkillsBlockLayout from './MyProfileEditPageLanguagesSkillsBlockLayout.svelte'
  import LabeledInputError from '@/common-components/LabeledInput/components/LabeledInputError/LabeledInputError.svelte'
  import Button from '@/common-components/Button/Button.svelte'
  import MyProfileEditPageLanguagesSkillsItemList from '../MyProfileEditPageLanguagesSkillsItemList/MyProfileEditPageLanguagesSkillsItemList.svelte'
  import { handleAddLanguagesSkillsButtonClick } from '../../handlers'

  type MyProfileEditPageLanguagesSkillsBlockProps = {
    formErrorsStore: Writable<Record<string, string[]>>
    formDataStore: Writable<MyProfileEditPageFormData>
    filterParams: FilterParams
    type: 'languages' | 'skills'
  }

  const { type, formErrorsStore, ...rest }: MyProfileEditPageLanguagesSkillsBlockProps = $props()

  const click = () =>
    handleAddLanguagesSkillsButtonClick(rest.formDataStore, rest.filterParams, type)
</script>

<MyProfileEditPageLanguagesSkillsBlockLayout>
  <Label text={type === 'languages' ? 'Spoken Languages' : 'Skills'} />
  <MyProfileEditPageLanguagesSkillsItemList {type} {...rest} />
  <LabeledInputError error={$formErrorsStore[type] && $formErrorsStore[type][0]} />
  <Button clickAction={click} styles="!mt-2 !w-40" color="grey"
    >Add {type === 'languages' ? 'A Language' : 'A Skill'}</Button
  >
</MyProfileEditPageLanguagesSkillsBlockLayout>
