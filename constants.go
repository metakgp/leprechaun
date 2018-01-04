package main

const ERP_SECRET_QUES_URL = "https://erp.iitkgp.ernet.in/SSOAdministration/getSecurityQues.htm"
const ERP_HOMEPAGE_URL = "https://erp.iitkgp.ernet.in/IIT_ERP3/welcome.jsp"

const PATH_INDEX_PAGE = "html/index.html"
const PATH_RESET_INDEX_PAGE = "html/reset_index.html"

const PATH_BEGIN_AUTH_PAGE = "templates/begin_auth.tmpl.html"
const PATH_BEGIN_RESET_PAGE = "templates/begin_reset.tmpl.html"
const PATH_BEGIN_AUTH_UNSUCCESSFUL_PAGE = "templates/begin_auth_unsuccessful.tmpl.html"
const PATH_STEP1_COMPLETE_PAGE = "templates/step1_complete.tmpl.html"
const PATH_STEP2_COMPLETE_PAGE = "templates/step2_complete.tmpl.html"
const PATH_RESET_COMPLETE_PAGE = "templates/reset_complete.tmpl.html"

const EMAIL_SUBJECT_STEP2 = "Leprechaun Authentication, Step 2 - Email Verification"
const EMAIL_SUBJECT_RESET = "Leprechaun Reset - Verification"

const ERROR_UNAUTH = "OOPS! You are not authenticated. If you would like to use Leprechaun, contact the Metakgp Maintainers. More info at https://wiki.metakgp.org/w/Metakgp:Governance#Current_maintainers"
