<script lang="ts">
    import {getFileServerHost} from "@/utils/utils.ts";

    const host = getFileServerHost();
    const link = host + '/files/home-banner.jpg';
</script>

<div class="absolute inset-0 bg-cover bg-center opacity-30"
     style="background-image: url({link});"
></div>
