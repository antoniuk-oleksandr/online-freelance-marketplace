<script lang="ts">
  import SearchSelectTriggerIcon from '../SearchSelectTriggerIcon/SearchSelectTriggerIcon.svelte'
  import SearchSelectTriggerLayout from './SearchSelectTriggerLayout.svelte'

  type SearchSelectTriggerProps = {
    selectedAdditionalItem: string
    selectedItem: string
  }

  const { selectedAdditionalItem, selectedItem }: SearchSelectTriggerProps = $props()
</script>

<SearchSelectTriggerLayout>
  <SearchSelectTriggerIcon {selectedAdditionalItem} />
  <span>{selectedItem}</span>
</SearchSelectTriggerLayout>
