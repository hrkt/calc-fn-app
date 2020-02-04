import io
import json
from decimal import *

from fdk import response


def handler(ctx, data: io.BytesIO=None):
    value = Decimal('NaN')
    try:
        body = json.loads(data.getvalue())
        left = Decimal(body.get('left'))
        right = Decimal(body.get('right'))
        value = left * right
    except (Exception, ValueError) as ex:
        print(str(ex))

    return response.Response(
        ctx, response_data=json.dumps(
            {"result": str(value)}),
        headers={"Content-Type": "application/json"}
    )
