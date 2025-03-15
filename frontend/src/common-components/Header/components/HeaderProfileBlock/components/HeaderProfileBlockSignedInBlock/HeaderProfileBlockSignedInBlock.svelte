<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import Dropdown from '@/common-components/Dropdown/Dropdown.svelte'
  import { getMyProfileDropdownItems } from '@/common-components/Header/helpers'
  import type { DropdownItem } from '@/types/DropdownItem'
  import type { SignHeaderData } from '@/types/SignHeaderData'

  type HeaderProfileBlockSignedInBlockProps = {
    signData: SignHeaderData
    darkMode: boolean | null
  }

  const { signData, darkMode }: HeaderProfileBlockSignedInBlockProps = $props()

  const dropdownItems: DropdownItem[] = $derived(getMyProfileDropdownItems(signData, darkMode))
</script>

<Dropdown menuWidth="min-w-header-dropdown-menu" modalHeaderTitle={'Menu'} items={dropdownItems}>
  <Avatar size="12" alt="user avatar" src={signData.avatar} />
</Dropdown>
