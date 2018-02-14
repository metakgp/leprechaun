package main

const ERP_SECRET_QUES_URL = "https://erp.iitkgp.ernet.in/SSOAdministration/getSecurityQues.htm"
const ERP_HOMEPAGE_URL = "https://erp.iitkgp.ernet.in/IIT_ERP3/welcome.jsp"

const PATH_INDEX_PAGE = "dist/html/index.html"
const PATH_RESET_INDEX_PAGE = "dist/html/reset_index.html"

const PATH_BEGIN_AUTH_PAGE = "dist/templates/begin_auth.tmpl.html"
const PATH_BEGIN_RESET_PAGE = "dist/templates/begin_reset.tmpl.html"
const PATH_BEGIN_AUTH_UNSUCCESSFUL_PAGE = "dist/templates/begin_auth_unsuccessful.tmpl.html"
const PATH_STEP1_COMPLETE_PAGE = "dist/templates/step1_complete.tmpl.html"
const PATH_STEP2_COMPLETE_PAGE = "dist/templates/step2_complete.tmpl.html"
const PATH_RESET_COMPLETE_PAGE = "dist/templates/reset_complete.tmpl.html"

const EMAIL_SUBJECT_STEP2 = "Leprechaun Authentication, Step 2 - Email Verification"
const EMAIL_SUBJECT_RESET = "Leprechaun Reset - Verification"

const ERROR_UNAUTH = "OOPS! You are not authenticated. If you would like to use Leprechaun, contact the Metakgp Maintainers. More info at https://wiki.metakgp.org/w/Metakgp:Governance#Current_maintainers"
const ERROR_UNKNOWN_TYPE = "That input type is not allowed! Allowed input types: roll, email"
const ERROR_NOT_FOUND = "That record doesn't exist in our database!"

const HASH_LEN = 15
