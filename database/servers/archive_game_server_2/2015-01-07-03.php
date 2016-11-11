<?php

$this->AddSQL("
update player_mail set expire_time = 0 where expire_time=1 and have_attachment=0;
");

