"use strict"
module.exports = function(settings) {
  var express = require("express");
  var router = express.Router();
  var logger = require("../log/logger");
  var HttpStatus = require("http-status-codes");
  var scalingEngineUtil = require("../utils/scalingEngineUtils")(settings.scalingEngine);
  var scalingHistoryHelper = require("./scalingHistoryHelper");

  router.get("/:app_id/scaling_histories", function(req, resp) {
    var appId = req.params.app_id;
    var startTime = req.query["start-time"];
    var endTime = req.query["end-time"];
    var order = req.query["order"];
    var page = req.query["page"];
    var resultsPerPage = req.query["results-per-page"];
    logger.info("Get scalinghistory", { "app_id": appId, "start-time": startTime, "end-time": endTime, "order": order, "page": page, "results-per-page": resultsPerPage });
    var parseResult = scalingHistoryHelper.parseParameter(req);
    if (!parseResult.valid) {
      logger.error("Failed to get scaling history", { "app_id": appId, "start-time": startTime, "end-time": endTime, "order": order, "page": page, "results-per-page": resultsPerPage, "message": parseResult.message });
      resp.status(HttpStatus.BAD_REQUEST).json({ "description": parseResult.message });
      return;
    }
    var parameters = parseResult.parameters;
    scalingEngineUtil.getScalingHistory(parameters, function(err, result) {
      var responseBody = {};
      var statusCode;
      if (err) {
        statusCode = err.statusCode;
        responseBody.description = err.message;
      } else {
        statusCode = result.statusCode;
        if (statusCode === HttpStatus.OK) {
          var page = parameters.page;
          var resultsPerPage = parameters.resultsPerPage;
          responseBody = scalingHistoryHelper.pagination(result.body, page, resultsPerPage);
        } else {
          responseBody.description = result.message;
        }
      }
      resp.status(statusCode).json(responseBody);
    });
  });
  return router;
}
