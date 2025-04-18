<script lang="ts">
  import { getFileServerHost } from '@/utils/utils'

  const host = getFileServerHost()
  const link = host + '/home-banner.jpg'
</script>

<div
  class="absolute inset-0 bg-cover bg-center opacity-30"
  style="background-image: url({link});"
></div>
