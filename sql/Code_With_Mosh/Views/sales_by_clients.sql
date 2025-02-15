
CREATE OR REPLACE VIEW sql_invoicing.sales_by_client AS
SELECT
	c.client_id,
    c.name,
    SUM(invoice_total) AS total_sales
FROM sql_invoicing.clients c
JOIN sql_invoicing.invoices i USING (client_id)
GROUP BY client_id, name;

CREATE OR REPLACE VIEW sql_invoicing.clients_balance AS
SELECT
	c.client_id,
    c.name,
    SUM(invoice_total - payment_total) AS balance
FROM sql_invoicing.clients c
JOIN sql_invoicing.invoices USING (client_id)
GROUP BY client_id, name
WITH CHECK OPTION;
