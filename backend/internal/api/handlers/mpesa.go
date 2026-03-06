package handlers

import "github.com/gin-gonic/gin"

// MpesaCallback handles incoming M-Pesa callback webhooks from Safaricom.
func MpesaCallback(c *gin.Context) {}

// MpesaSTKPush initiates an M-Pesa STK push (authenticated action).
func MpesaSTKPush(c *gin.Context) {}

// MpesaStatus returns the status of an M-Pesa transaction.
func MpesaStatus(c *gin.Context) {}
