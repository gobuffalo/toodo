require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

window.watchCheckboxes = function() {
  $(".item-form input[type=checkbox]").on("change", (e) => {
    let $e = $(e.target);
    let $form = $e.closest("form");
    $form.submit();
  });
};

$(() => {
  watchCheckboxes();
});
