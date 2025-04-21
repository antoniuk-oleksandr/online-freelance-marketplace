<script lang="ts">
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import type { ChatPartner } from '@/types/ChatPartner'
  import OrderByIdChatTabHeaderLayout from './OrderByIdChatTabHeaderLayout.svelte'
  import { makeLastOnlineText } from '../../helpers'
  import { Link } from 'svelte-routing'

  type OrderByIdChatTabHeaderProps = ChatPartner & {}

  const { id, avatar, username, lastOnline }: OrderByIdChatTabHeaderProps = $props()
</script>

<OrderByIdChatTabHeaderLayout>
  <Avatar src={avatar} alt="avatar" />
  <div class="flex flex-col">
    <Link
      class="font-semibold hover:underline text-light-palette-text-primary dark:text-dark-palette-text-primary"
      to="/users/{id}"
    >
      {username}
    </Link>
    <span class="text-cyan-500">{makeLastOnlineText(lastOnline)}</span>
  </div>
</OrderByIdChatTabHeaderLayout>
