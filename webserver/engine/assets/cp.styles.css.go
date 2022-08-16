package assets

var CpStylesCss = []byte(`
/*
select.bs-select-hidden,.bootstrap-select > select.bs-select-hidden,select.selectpicker {
  display: none!important;
}

.bootstrap-select {
  width: 220px \0;
  vertical-align: middle;
}

.bootstrap-select > .dropdown-toggle {
  position: relative;
  width: 100%;
  text-align: right;
  white-space: nowrap;
  display: -webkit-inline-box;
  display: -webkit-inline-flex;
  display: -ms-inline-flexbox;
  display: inline-flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
  -webkit-box-pack: justify;
  -webkit-justify-content: space-between;
  -ms-flex-pack: justify;
  justify-content: space-between;
}

.bootstrap-select > .dropdown-toggle:after {
  margin-top: -1px;
}

.bootstrap-select > .dropdown-toggle.bs-placeholder,.bootstrap-select > .dropdown-toggle.bs-placeholder:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder:active {
  color: #999;
}

.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-primary,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-secondary,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-success,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-danger,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-info,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-dark,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-primary:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-secondary:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-success:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-danger:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-info:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-dark:hover,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-primary:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-secondary:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-success:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-danger:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-info:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-dark:focus,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-primary:active,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-secondary:active,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-success:active,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-danger:active,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-info:active,.bootstrap-select > .dropdown-toggle.bs-placeholder.btn-dark:active {
  color: rgba(255,255,255,0.5);
}

.bootstrap-select > select {
  position: absolute!important;
  bottom: 0;
  left: 50%;
  display: block!important;
  width: .5px!important;
  height: 100%!important;
  padding: 0!important;
  opacity: 0!important;
  border: none;
  z-index: 0!important;
}

.bootstrap-select > select.mobile-device {
  top: 0;
  left: 0;
  display: block!important;
  width: 100%!important;
  z-index: 2!important;
}

.has-error .bootstrap-select .dropdown-toggle,.error .bootstrap-select .dropdown-toggle,.bootstrap-select.is-invalid .dropdown-toggle,.was-validated .bootstrap-select .selectpicker:invalid + .dropdown-toggle {
  border-color: #b94a48;
}

.bootstrap-select.is-valid .dropdown-toggle,.was-validated .bootstrap-select .selectpicker:valid + .dropdown-toggle {
  border-color: #28a745;
}

.bootstrap-select.fit-width {
  width: auto!important;
}

.bootstrap-select:not([class*="col-"]):not([class*="form-control"]):not(.input-group-btn) {
  width: 220px;
}

.bootstrap-select > select.mobile-device:focus + .dropdown-toggle,.bootstrap-select .dropdown-toggle:focus {
  outline: thin dotted #333!important;
  outline: 5px auto -webkit-focus-ring-color!important;
  outline-offset: -2px;
}

.bootstrap-select.form-control {
  margin-bottom: 0;
  padding: 0;
  border: none;
}

:not(.input-group) > .bootstrap-select.form-control:not([class*="col-"]) {
  width: 100%;
}

.bootstrap-select.form-control.input-group-btn {
  float: none;
  z-index: auto;
}

.form-inline .bootstrap-select,.form-inline .bootstrap-select.form-control:not([class*="col-"]) {
  width: auto;
}

.bootstrap-select:not(.input-group-btn),.bootstrap-select[class*="col-"] {
  float: none;
  display: inline-block;
  margin-left: 0;
}

.bootstrap-select.dropdown-menu-right,.bootstrap-select[class*="col-"].dropdown-menu-right,.row .bootstrap-select[class*="col-"].dropdown-menu-right {
  float: right;
}

.form-inline .bootstrap-select,.form-horizontal .bootstrap-select,.form-group .bootstrap-select {
  margin-bottom: 0;
}

.form-group-lg .bootstrap-select.form-control,.form-group-sm .bootstrap-select.form-control {
  padding: 0;
}

.form-group-lg .bootstrap-select.form-control .dropdown-toggle,.form-group-sm .bootstrap-select.form-control .dropdown-toggle {
  height: 100%;
  font-size: inherit;
  line-height: inherit;
  border-radius: inherit;
}

.bootstrap-select.form-control-sm .dropdown-toggle,.bootstrap-select.form-control-lg .dropdown-toggle {
  font-size: inherit;
  line-height: inherit;
  border-radius: inherit;
}

.bootstrap-select.form-control-sm .dropdown-toggle {
  padding: .25rem .5rem;
}

.bootstrap-select.form-control-lg .dropdown-toggle {
  padding: .5rem 1rem;
}

.form-inline .bootstrap-select .form-control {
  width: 100%;
}

.bootstrap-select.disabled,.bootstrap-select > .disabled {
  cursor: not-allowed;
}

.bootstrap-select.disabled:focus,.bootstrap-select > .disabled:focus {
  outline: none!important;
}

.bootstrap-select.bs-container {
  position: absolute;
  top: 0;
  left: 0;
  height: 0!important;
  padding: 0!important;
}

.bootstrap-select.bs-container .dropdown-menu {
  z-index: 1060;
}

.bootstrap-select .dropdown-toggle .filter-option {
  position: static;
  top: 0;
  left: 0;
  float: left;
  height: 100%;
  width: 100%;
  text-align: left;
  overflow: hidden;
  -webkit-box-flex: 0;
  -webkit-flex: 0 1 auto;
  -ms-flex: 0 1 auto;
  flex: 0 1 auto;
}


.bs3.bootstrap-select .dropdown-toggle .filter-option {
  padding-right: inherit;
}

.input-group .bs3-has-addon.bootstrap-select .dropdown-toggle .filter-option {
  position: absolute;
  padding-top: inherit;
  padding-bottom: inherit;
  padding-left: inherit;
  float: none;
}

.input-group .bs3-has-addon.bootstrap-select .dropdown-toggle .filter-option .filter-option-inner {
  padding-right: inherit;
}

.bootstrap-select .dropdown-toggle .filter-option-inner-inner {
  overflow: hidden;
}

.bootstrap-select .dropdown-toggle .filter-expand {
  width: 0!important;
  float: left;
  opacity: 0!important;
  overflow: hidden;
}

.bootstrap-select .dropdown-toggle .caret {
  position: absolute;
  top: 50%;
  right: 12px;
  margin-top: -2px;
  vertical-align: middle;
}

.input-group .bootstrap-select.form-control .dropdown-toggle {
  border-radius: inherit;
}

.bootstrap-select[class*="col-"] .dropdown-toggle {
  width: 100%;
}

.bootstrap-select .dropdown-menu {
  min-width: 100%;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

.bootstrap-select .dropdown-menu > .inner:focus {
  outline: none!important;
}

.bootstrap-select .dropdown-menu.inner {
  position: static;
  float: none;
  border: 0;
  padding: 0;
  margin: 0;
  border-radius: 0;
  -webkit-box-shadow: none;
  box-shadow: none;
}

.bootstrap-select .dropdown-menu li {
  position: relative;
}

.bootstrap-select .dropdown-menu li.active small {
  color: rgba(255,255,255,0.5)!important;
}

.bootstrap-select .dropdown-menu li.disabled a {
  cursor: not-allowed;
}

.bootstrap-select .dropdown-menu li a {
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.bootstrap-select .dropdown-menu li a.opt {
  position: relative;
  padding-left: 2.25em;
}

.bootstrap-select .dropdown-menu li a span.check-mark {
  display: none;
}

.bootstrap-select .dropdown-menu li a span.text {
  display: inline-block;
}

.bootstrap-select .dropdown-menu li small {
  padding-left: .5em;
}

.bootstrap-select .dropdown-menu .notify {
  position: absolute;
  bottom: 5px;
  width: 96%;
  margin: 0 2%;
  min-height: 26px;
  padding: 3px 5px;
  background: #f5f5f5;
  border: 1px solid #e3e3e3;
  -webkit-box-shadow: inset 0 1px 1px rgba(0,0,0,0.05);
  box-shadow: inset 0 1px 1px rgba(0,0,0,0.05);
  pointer-events: none;
  opacity: .9;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

.bootstrap-select .no-results {
  padding: 3px;
  background: #f5f5f5;
  margin: 0 5px;
  white-space: nowrap;
}

.bootstrap-select.fit-width .dropdown-toggle .filter-option {
  position: static;
  display: inline;
  padding: 0;
  width: auto;
}

.bootstrap-select.fit-width .dropdown-toggle .filter-option-inner,.bootstrap-select.fit-width .dropdown-toggle .filter-option-inner-inner {
  display: inline;
}

.bootstrap-select.fit-width .dropdown-toggle .bs-caret:before {
  content: '\00a0';
}

.bootstrap-select.fit-width .dropdown-toggle .caret {
  position: static;
  top: auto;
  margin-top: -1px;
}

.bootstrap-select.show-tick .dropdown-menu .selected span.check-mark {
  position: absolute;
  display: inline-block;
  right: 15px;
  top: 5px;
}

.bootstrap-select.show-tick .dropdown-menu li a span.text {
  margin-right: 34px;
}

.bootstrap-select .bs-ok-default:after {
  content: '';
  display: block;
  width: .5em;
  height: 1em;
  border-style: solid;
  border-width: 0 .26em .26em 0;
  -webkit-transform: rotate(45deg);
  -ms-transform: rotate(45deg);
  -o-transform: rotate(45deg);
  transform: rotate(45deg);
}

.bootstrap-select.show-menu-arrow.open > .dropdown-toggle,.bootstrap-select.show-menu-arrow.show > .dropdown-toggle {
  z-index: 1061;
}

.bootstrap-select.show-menu-arrow .dropdown-toggle .filter-option:before {
  content: '';
  border-left: 7px solid transparent;
  border-right: 7px solid transparent;
  border-bottom: 7px solid rgba(204,204,204,0.2);
  position: absolute;
  bottom: -4px;
  left: 9px;
  display: none;
}

.bootstrap-select.show-menu-arrow .dropdown-toggle .filter-option:after {
  content: '';
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-bottom: 6px solid #fff;
  position: absolute;
  bottom: -4px;
  left: 10px;
  display: none;
}

.bootstrap-select.show-menu-arrow.dropup .dropdown-toggle .filter-option:before {
  bottom: auto;
  top: -4px;
  border-top: 7px solid rgba(204,204,204,0.2);
  border-bottom: 0;
}

.bootstrap-select.show-menu-arrow.dropup .dropdown-toggle .filter-option:after {
  bottom: auto;
  top: -4px;
  border-top: 6px solid #fff;
  border-bottom: 0;
}

.bootstrap-select.show-menu-arrow.pull-right .dropdown-toggle .filter-option:before {
  right: 12px;
  left: auto;
}

.bootstrap-select.show-menu-arrow.pull-right .dropdown-toggle .filter-option:after {
  right: 13px;
  left: auto;
}

.bootstrap-select.show-menu-arrow.open > .dropdown-toggle .filter-option:before,.bootstrap-select.show-menu-arrow.show > .dropdown-toggle .filter-option:before,.bootstrap-select.show-menu-arrow.open > .dropdown-toggle .filter-option:after,.bootstrap-select.show-menu-arrow.show > .dropdown-toggle .filter-option:after {
  display: block;
}
*/

.bs-searchbox,.bs-actionsbox,.bs-donebutton {
  padding: 4px 8px;
}

.bs-actionsbox {
  width: 100%;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

.bs-actionsbox .btn-group button {
  width: 50%;
}

.bs-donebutton {
  float: left;
  width: 100%;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  box-sizing: border-box;
}

.bs-donebutton .btn-group button {
  width: 100%;
}

.bs-searchbox + .bs-actionsbox {
  padding: 0 8px 4px;
}

.bs-searchbox .form-control {
  margin-bottom: 0;
  width: 100%;
  float: none;
}

body.cp {
  display: block;
  padding-right: 0!important;
}

.cp .navbar {
  padding: .5rem 1rem!important;
}

.dropdown-item:focus,.dropdown-item:hover {
  background-color: #f1f1f1;
}

.dropdown-item.active,.dropdown-item:active {
  background-color: #007bff;
}

html {
  height: 100%;
}

body.cp-login,body.cp-mysql,body.cp-first-user {
  min-height: 100%;
  display: -ms-flexbox;
  display: -webkit-box;
  display: flex;
  -ms-flex-align: center;
  -ms-flex-pack: center;
  -webkit-box-align: center;
  align-items: center;
  -webkit-box-pack: center;
  justify-content: center;
  background-color: #eee;
}

.cp-login .form-signin,.cp-mysql .form-signin,.cp-first-user .form-signin {
  width: 100%;
  max-width: 21rem;
  padding: 1.5rem;
  margin: 1.5rem;
}

.cp-login .form-signin label,.cp-mysql .form-signin label,.cp-first-user .form-signin label {
  cursor: pointer;
}

.cp-login .form-signin .form-control,.cp-mysql .form-signin .form-control,.cp-first-user .form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: .5rem;
  font-size: 1rem;
}

.cp-login .form-signin .form-control:focus,.cp-mysql .form-signin .form-control:focus,.cp-first-user .form-signin .form-control:focus {
  z-index: 2;
}

.cp-login .form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.cp-login .form-signin input[type="password"] {
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}

.cp-login .sys-messages,.cp-mysql .sys-messages,.cp-first-user .sys-messages {
  text-align: left;
}

.cp-login .sys-back-link {
  padding-top: 1.5rem;
}

body.cp {
  font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol","Noto Color Emoji";
  background: initial;
  background-color: #fff;
  font-size: 1rem;
  font-weight: 400;
  line-height: 1.5;
  color: #444;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

body.cp nav.main {
  height: 3.5rem;
  box-shadow: 0 0 5px 4px rgba(0,0,0,0.25);
}

body.cp nav.main.bg-dark {
  background: #205081 url(/assets/sys/bg.png) repeat 0 0!important;
}

body.cp nav.main a.navbar-brand {
  font-weight: 700;
}

body.cp .wrap {
  width: 100%;
  height: 100%;
  display: block;
  align-items: stretch;
}

body.cp .wrap .sidebar,body.cp .wrap .content {
  display: block;
  height: 100%;
  position: relative;
  padding-top: 3.5rem;
  vertical-align: top;
}

body.cp .wrap .sidebar {
  float: left;
  z-index: 1;
}

/*
body.cp .wrap .content {
  display: block!important;
  z-index: 0;
}
*/

/*
body.cp .content {
  display: block!important;
  z-index: 0;
}
*/

body.cp .content .sidebar.sidebar-right {
  float: right;
  display: none!important;
  z-index: 1;
}

body.cp .content .sidebar-right .padd,
body.cp .content .padd {
  padding: 1rem;
}

body.cp .content .scroll {
  height: 100%;
  overflow: hidden;
  overflow-y: auto;
}

/*
body.cp .content .sidebar {
  width: 15.4rem;
  background: #eee;
  box-shadow: 0 .5em .5em rgba(0,0,0,.3);
}
*/

body.cp .wrap .sidebar .dropdown-divider {
  border-color: #d6d6d6;
  margin: 0;
}

body.cp .wrap .sidebar.sidebar-left ul.nav {
  padding: 1rem 0;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-separator {
  height: 1rem;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item a {
  color: #444;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.disabled a {
  color: #c9302c;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.active {
  background-color: #417cb9;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.active.disabled {
  background-color: #c9302c;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.active a {
  color: #fff;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item:hover {
  background-color: #e7e7e7;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.active:hover {
  background-color: #417cb9;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.active.disabled:hover {
  background-color: #c9302c;
}

body.cp .wrap .sidebar.sidebar-left ul.nav ul.nav {
  background: #eee;
  padding-top: 0;
}

body.cp .wrap .sidebar.sidebar-left ul.nav ul.nav li.nav-item a {
  color: #444;
  padding-left: 2rem;
}

body.cp .wrap .sidebar.sidebar-left ul.nav ul.nav li.nav-item.active {
  background-color: #e7e7e7;
}

body.cp .wrap .sidebar.sidebar-left ul.nav ul.nav li.nav-item.active a {
  color: #417cb9;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item.disabled ul.nav li.nav-item.active a {
  color: #c9302c;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item:last-child ul {
  padding-bottom: 0;
}

body.cp .wrap .sidebar.sidebar-left ul.nav li.nav-item svg.sicon {
  fill: currentColor;
  margin-right: .3rem;
  margin-bottom: 1px;
}

.svg-green svg {
  fill: currentColor;
  color: #28a745;
}

.svg-red svg {
  fill: currentColor;
  color: #d9534f;
}

.svg-yellow svg {
  fill: currentColor;
  color: #f0ad4e;
}

ul.pagination {
  margin-top: 1rem;
}

.data-table {
  margin-bottom: 0;
}

.data-table.table-hover tbody tr:hover {
  background-color: #fffbdf;
}

.data-table a svg {
  fill: currentColor;
  color: #007bff;
}

.data-table a:hover svg {
  color: #0056b3;
}

.data-table a.ico.delete svg {
  color: #d9534f;
}

.data-table a.ico.delete:hover svg {
  color: #c9302c;
}

.data-table td.col_action a.ico {
  display: inline-block;
  width: 1rem;
  height: 1rem;
  margin-right: .6rem;
}

.data-table td.col_action a.ico:last-child {
  margin-right: 0;
}

.data-table thead tr {
  background-color: #e9ecef;
}

.data-table.table-bordered td,.data-table.table-bordered th {
  border: none;
  border-top: 1px solid #dee2e6;
}

.data-table.table-bordered td .parent,.data-table.table-bordered td .parent a,.data-table.table-bordered td .template {
  color: #adafb2;
}

.data-table.table_pages .col_datetime {
  width: 8rem;
}

.data-table.table_pages .col_active {
  width: 5rem;
}

.data-table.table_pages .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_blog_posts .col_datetime {
  width: 8rem;
}

.data-table.table_blog_posts .col_active {
  width: 5rem;
}

.data-table.table_blog_posts .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_blog_cats .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_shop_products .col_price {
  width: 8rem;
}

.data-table.table_shop_products .col_datetime {
  width: 8rem;
}

.data-table.table_shop_products .col_active {
  width: 5rem;
}

.data-table.table_shop_products .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_shop_products .col_quantity {
  width: 8rem;
}

.data-table.table_shop_products .col_total {
  width: 8rem;
}

.data-table.table_shop_cats .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_shop_filters .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_shop_currencies .col_coefficient {
  width: 7rem;
}

.data-table.table_shop_currencies .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_users .col_active,.data-table.table_users .col_admin {
  width: 5rem;
}

.data-table.table_users .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_notify_mail .col_datetime {
  width: 8rem;
}

.data-table.table_notify_mail .col_status {
  width: 5rem;
}

.data-table.table_shop_orders .col_id {
  width: 7rem;
}

.data-table.table_shop_orders .col_create_datetime {
  width: 8rem;
}

.data-table.table_shop_orders .col_client_email {
  width: 8rem;
}

/*
.data-form label {
  font-weight: 700;
  margin-top: .45rem;
  margin-bottom: .45rem;
}
*/

.data-form small {
  color: #aeb8bc;
}

.data-form > div:nth-last-child(2) {
  margin-bottom: 0;
}

.data-form textarea {
  min-height: 5.4rem;
}

.checkbox-ios {
  display: inline-block;
}

.checkbox-ios input[type=checkbox] {
  max-height: 0;
  max-width: 0;
  opacity: 0;
  position: absolute;
}

.checkbox-ios input[type=checkbox] + label {
  display: block;
  position: relative;
  box-shadow: inset 0 0 0 1px #ced4da;
  background: #ced4da;
  text-indent: -5000px;
  height: 30px;
  width: 60px;
  border-radius: 1.5rem;
  cursor: pointer;
  margin-top: 0;
  margin-bottom: 0;
}

.checkbox-ios input[type=checkbox] + label:before {
  content: "";
  position: absolute;
  display: block;
  height: 30px;
  width: 30px;
  top: 0;
  left: 0;
  border-radius: 1.5rem;
  background: rgba(19,191,17,0);
  -moz-transition: .1s ease-in-out;
  -webkit-transition: .1s ease-in-out;
  transition: .1s ease-in-out;
}

.checkbox-ios input[type=checkbox] + label:after {
  content: "";
  position: absolute;
  display: block;
  height: 26px;
  width: 26px;
  top: 2px;
  left: 2px;
  border-radius: 1.5rem;
  background: #fff;
  -moz-transition: .1s ease-in-out;
  -webkit-transition: .1s ease-in-out;
  transition: .1s ease-in-out;
}

.checkbox-ios input[type=checkbox]:checked + label:before {
  width: 60px;
  background: #1a73e8;
}

.checkbox-ios input[type=checkbox]:checked + label:after {
  left: 32px;
}

.checkbox-clickable {
  display: inline-block;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  margin-right: .5rem;
}

.checkbox-clickable:last-child {
  margin-right: 0;
}

.checkbox-clickable input[type=checkbox] {
  max-height: 0;
  max-width: 0;
  opacity: 0;
  position: absolute;
}

.checkbox-clickable input[type=checkbox] + label {
  display: block;
  position: relative;
  box-shadow: inset 0 0 0 1px #ced4da;
  background: #e9ecef;
  height: 28px;
  border-radius: 1.5rem;
  cursor: pointer;
  margin-top: .1rem;
  margin-bottom: .1rem;
  padding: 0 1rem;
  line-height: 1.8rem;
  font-size: 14px;
  color: #abb3bb;
}

.checkbox-clickable input[type=checkbox]:checked + label {
  box-shadow: inset 0 0 0 1px #1a73e8;
  background: #1a73e8;
  color: #fff;
}

.list-wrapper {
  background: #e9ecef;
  padding: 1rem;
  border-radius: .25rem;
}

.btn-dynamic-remove {
  position: absolute;
  top: 0;
  right: 0;
}

.product-copy,.product-another {
  position: relative;
}

.product-copy a,.product-another a {
  position: absolute;
  right: 0;
  padding: 12px 18px;
  background: #e9ecef;
  color: #6c757d;
  border-radius: 0 .25rem .25rem 0;
}

.product-another a {
  right: 52px;
}

.product-copy a svg,.product-another a svg {
  fill: currentColor;
}

.product-copy a:hover,.product-another a:hover {
  color: #417cb9;
}

#list-images {
  display: block;
}

#list-images .attached-img {
  display: inline-block;
  padding: 1rem;
  background: #fff;
  margin-right: 1rem;
  margin-bottom: 1rem;
  border-radius: .25rem;
  position: relative;
  vertical-align: top;
}

#list-images .attached-img a {
  display: block;
}

#list-images .attached-img a img {
  width: 84px;
  height: 84px;
}

#list-images .attached-img a.remove {
  position: absolute;
  width: 24px;
  height: 24px;
  top: -12px;
  right: -12px;
  border-radius: 12px;
  background: #d9534f;
  text-align: center;
  cursor: pointer;
}

#list-images .attached-img a.remove:hover {
  background: #c9302c;
}

#list-images .attached-img a.remove svg {
  color: #fff;
  fill: currentColor;
}

#list-images .attached-img.sortable-ghost {
  opacity: .1;
}

#list-images.loading .attached-img {
  opacity: .1;
}

#list-images.loading .attached-img a.remove {
  display: none;
}

#upload-msg {
  position: absolute;
  background: #fff;
  width: 100%;
  height: 100%;
  border: 1px solid #ced4da;
  border-radius: .25rem;
  padding: .375rem .75rem;
  display: none;
}

.data-table.table_fm_files .col_name {
  max-width: 10rem;
}

.data-table.table_fm_files .col_type {
  width: 8rem;
}

.data-table.table_fm_files .col_action {
  width: 6rem;
  text-align: right;
}

.data-table.table_fm_files .col_name a {
  display: block;
}

#sys-modal-files-manager #fmfiles {
  display: block;
}

#sys-modal-files-manager button.upload {
  display: none;
  width: 100%;
}

#sys-modal-files-manager.uploading #fmfiles {
  display: none;
}

#sys-modal-files-manager.uploading button.upload {
  display: block;
}

.fm-loading {
  min-width: 99px;
  min-height: 99px;
  background: #fff url(/assets/cp/img-load.gif) no-repeat 50% 50%!important;
}

/*
.dropdown.bootstrap-select {
  position: relative;
}

.dropdown.bootstrap-select select {
  position: static!important;
}

.dropdown.bootstrap-select button.dropdown-toggle {
  position: absolute;
  top: 0;
  left: 0;
}
*/

#sys-modal-user-settings {
  padding-right: 0!important;
}

textarea.form-control.wysiwyg {
  min-height: 340px;
}

div.wysiwyg.focused {
  background-color: #fff;
  border-color: #80bdff;
  outline: 0;
  box-shadow: 0 0 0 .2rem rgba(0,123,255,.25);
}

.CodeMirror {
  border: 1px solid #eee;
  height: auto;
}

.CodeMirror pre > * {
  text-indent: 0;
}

.btn-primary {
  background-color: #1a73e8;
  border-color: #1a73e8;
  color: #fff;
}

.btn-primary:hover {
  background: rgba(26,115,232,0.761);
}

.btn-secondary {
  background: transparent;
  border-color: #dadce0;
  color: #1a73e8;
  font-weight: 500;
}

.btn-secondary:hover {
  background: rgba(66,133,244,0.04);
  border-color: #d2e3fc;
  color: #1a73e8;
}

.bs-searchbox,.bs-actionsbox,.bs-donebutton {
  padding: 0 8px 8px;
}

/*
.bootstrap-select button.btn {
  outline: 0;
  border-color: #ced4da;
}

.bootstrap-select button.btn:active,.bootstrap-select button.btn:hover,.bootstrap-select button.btn:focus {
  outline: 0!important;
}

.bootstrap-select ul.dropdown-menu li {
  background-color: transparent;
}

.bootstrap-select.form-control.show {
  box-shadow: 0 0 0 .2rem rgba(0,123,255,.25);
}

.bootstrap-select.form-control.show button.btn {
  border-color: #80bdff;
  background-color: #fff;
}

.bootstrap-select .dropdown-menu li a.selected {
  background-color: #007bff;
  color: #fff;
}

.bootstrap-select .dropdown-menu.show {
  padding-bottom: 0!important;
}

*/

.text-dotted {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  display: inline-block;
  max-width: 100%;
  margin-bottom: -.5rem;
}

@media (min-width: 992px) {
  body.cp.cp-sidebar-right .wrap .sidebar.sidebar-right.d-lg-table-cell {
    display: table-cell!important;
  }

  .data-form .row .sys-messages .alert {
    margin-top: 1rem;
    margin-bottom: 0;
  }

  .data-form .form-group.last {
    margin-bottom: 0;
  }

  .data-form.index-add .form-group.n9,.data-form.index-modify .form-group.n9 {
    margin-bottom: 0;
  }

  .data-form.blog-add .form-group.n8,.data-form.blog-modify .form-group.n8 {
    margin-bottom: 0;
  }

  .data-form.blog-categories-add .form-group.n4,.data-form.blog-categories-modify .form-group.n4 {
    margin-bottom: 0;
  }

  .data-form.shop-add .form-group.n15,.data-form.shop-modify .form-group.n15 {
    margin-bottom: 0;
  }

  .data-form.shop-categories-add .form-group.n4,.data-form.shop-categories-modify .form-group.n4 {
    margin-bottom: 0;
  }

  .data-form.shop-attributes-add .form-group.n4,.data-form.shop-attributes-modify .form-group.n4 {
    margin-bottom: 0;
  }

  .data-form.shop-currencies-add .form-group.n5,.data-form.shop-currencies-modify .form-group.n5 {
    margin-bottom: 0;
  }

  .data-form.users-add .form-group.n7,.data-form.users-modify .form-group.n7 {
    margin-bottom: 0;
  }

  .data-form.settings-pagination .form-group.n2 {
    margin-bottom: 0;
  }
}

@media (max-width: 575px) {
  body.cp {
    height: auto;
    overflow: none;
    overflow-y: scroll;
  }

  body.cp .wrap .sidebar {
    width: auto;
    float: none;
    box-shadow: none;
  }

  body.cp .wrap .content {
    padding-top: 0;
  }
}

@media (max-width: 767px) {
  .navbar-expand-md .navbar-collapse {
    padding: 1rem;
    background: #417cb9;
    box-shadow: 0 .2em .2em rgba(0,0,0,.3);
    border-radius: .25rem;
  }
}`)
