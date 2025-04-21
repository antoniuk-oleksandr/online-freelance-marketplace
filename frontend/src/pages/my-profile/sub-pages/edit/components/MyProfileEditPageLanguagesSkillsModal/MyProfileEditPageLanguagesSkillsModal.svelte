<script module lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import NewSelect from '@/common-components/NewSelect/NewSelect.svelte'
  import type { FilterParams } from '@/types/FilterParams'
  import type { MyProfileEditPageFormData } from '@/types/MyProfileEditPageFormData'
  import { get, type Writable } from 'svelte/store'
  import { extractNewSelectItemsFromFilterParams } from '../../helpers'
  import { handleLanguagesSkillsModalSubmit } from '../../handlers'

  type MyProfileEditPageLanguagesSkillsModalProps = {
    filterParams: FilterParams
    type: 'languages' | 'skills'
    formDataStore: Writable<MyProfileEditPageFormData>
  }

  let selectedValue = $state<string>('')
  const setSelectedValue = (value: string) => (selectedValue = value)

  export { myProfileEditPageLanguagesSkillsModal }
</script>

{#snippet myProfileEditPageLanguagesSkillsModal(props: MyProfileEditPageLanguagesSkillsModalProps)}
  <div class="flex flex-col gap-6 p-6">
    <NewSelect
      noneSelectedText="Select a {props.type.slice(0, -1)}"
      selectWidth="w-full"
      onSelectChange={(value) => setSelectedValue(value as string)}
      initialSelectValues={extractNewSelectItemsFromFilterParams(
        props.filterParams,
        props.type,
        get(props.formDataStore),
      )}
      modalHeaderTitle="Select a {props.type.slice(0, -1)}"
      initialSelectIndex={-1}
    />
    <Button
      clickAction={() =>
        handleLanguagesSkillsModalSubmit(
          selectedValue,
          props.type,
          props.formDataStore,
          props.filterParams,
        )}
      type="button">Submit</Button
    >
  </div>
{/snippet}
